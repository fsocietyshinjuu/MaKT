// Mouse and Keyboard Test launcher.
// Serves the embedded page on 127.0.0.1 and opens it as a standalone window
// in an installed Chromium-based browser (Edge, Chrome, Brave or Chromium).
package main

import (
	"bytes"
	_ "embed"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync/atomic"
	"time"
)

//go:embed index.html
var page []byte

// Only used when no Chromium browser is found and the page opens in the default browser:
// the page pings the launcher so the launcher can quit after the tab is closed.
const heartbeat = `<script>setInterval(function(){fetch('/ping',{cache:'no-store'}).catch(function(){})},4000)</script></body>`

func browserCandidates() []string {
	var c []string
	switch runtime.GOOS {
	case "windows":
		for _, env := range []string{"ProgramFiles", "ProgramFiles(x86)", "LocalAppData"} {
			base := os.Getenv(env)
			if base == "" {
				continue
			}
			for _, rel := range []string{
				`Microsoft\Edge\Application\msedge.exe`,
				`Google\Chrome\Application\chrome.exe`,
				`BraveSoftware\Brave-Browser\Application\brave.exe`,
				`Chromium\Application\chrome.exe`,
			} {
				c = append(c, filepath.Join(base, rel))
			}
		}
	case "darwin":
		home, _ := os.UserHomeDir()
		for _, root := range []string{"/Applications", filepath.Join(home, "Applications")} {
			for _, app := range [][2]string{
				{"Google Chrome", "Google Chrome"},
				{"Microsoft Edge", "Microsoft Edge"},
				{"Brave Browser", "Brave Browser"},
				{"Chromium", "Chromium"},
			} {
				c = append(c, filepath.Join(root, app[0]+".app", "Contents", "MacOS", app[1]))
			}
		}
	default:
		for _, name := range []string{
			"google-chrome", "google-chrome-stable", "chromium", "chromium-browser",
			"microsoft-edge", "microsoft-edge-stable", "brave-browser", "brave",
		} {
			if p, err := exec.LookPath(name); err == nil {
				c = append(c, p)
			}
		}
		c = append(c, "/snap/bin/chromium")
	}
	return c
}

func findBrowser() string {
	if p := os.Getenv("MKT_BROWSER"); p != "" {
		return p
	}
	for _, p := range browserCandidates() {
		if st, err := os.Stat(p); err == nil && !st.IsDir() {
			return p
		}
	}
	return ""
}

func openDefault(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	_ = cmd.Start()
}

func main() {
	var lastPing atomic.Int64
	lastPing.Store(time.Now().UnixNano())

	withBeat := bytes.Replace(page, []byte("</body>"), []byte(heartbeat), 1)
	var withBeatServed atomic.Bool

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		lastPing.Store(time.Now().UnixNano())
		w.WriteHeader(http.StatusNoContent)
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Cache-Control", "no-store")
		if withBeatServed.Load() {
			_, _ = w.Write(withBeat)
		} else {
			_, _ = w.Write(page)
		}
	})

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		os.Exit(1)
	}
	go func() { _ = http.Serve(ln, mux) }()
	url := "http://" + ln.Addr().String() + "/"

	if b := findBrowser(); b != "" {
		profile, err := os.MkdirTemp("", "mouse-keyboard-test-")
		if err == nil {
			defer os.RemoveAll(profile)
		}
		cmd := exec.Command(b,
			"--app="+url,
			"--user-data-dir="+profile,
			"--no-first-run",
			"--no-default-browser-check",
			"--window-size=1280,860",
		)
		if err := cmd.Start(); err == nil {
			_ = cmd.Wait()
			return
		}
	}

	// Fallback: default browser, quit when the tab stops pinging.
	withBeatServed.Store(true)
	openDefault(url)
	for {
		time.Sleep(2 * time.Second)
		if time.Since(time.Unix(0, lastPing.Load())) > 20*time.Second {
			return
		}
	}
}
