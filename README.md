# Mouse and Keyboard Test

A free mouse and keyboard tester. Press a button or a key and its picture lights up. It also warns about double presses, missed presses and a scroll wheel that jumps backwards, it can measure your mouse DPI, and it tests your speakers and microphone. Nothing leaves your computer.

- **Landing page:** `index.html`
- **Web app:** `app/index.html` (opens at `/app/`)
- **Desktop downloads:** `downloads/` (Windows, macOS, Linux) with `SHA256SUMS.txt`

## Structure

```
index.html            landing page
app/index.html        the web app
assets/               screenshots and favicon
downloads/            zipped desktop apps + checksums
launcher/             Go source and build script for the desktop apps
```

Everything is plain static files. There is no build step.

## Deploy on GitHub Pages

1. Create a new repository on GitHub (for example `mouse-keyboard-test`).
2. Upload the **contents** of this folder (`index.html`, `app`, `assets`, `downloads`, `README.md`) to the repository root.
3. Go to **Settings > Pages**.
4. Under **Build and deployment**, set **Source** to **Deploy from a branch**, choose branch **main** and folder **/ (root)**, then **Save**.
5. Wait a minute or two. The site appears at `https://YOUR-USERNAME.github.io/mouse-keyboard-test/`.

The web app will be at `https://YOUR-USERNAME.github.io/mouse-keyboard-test/app/`.

All links are relative, so the site also works from a custom domain or from a user site repository named `YOUR-USERNAME.github.io`.

## Preview locally

```
python3 -m http.server 8000
```

Then open http://localhost:8000.

## Updating the downloads

The desktop apps contain a copy of `app/index.html`, so rebuild them whenever the web app changes.

1. Install Go 1.22 or newer and run `sh launcher/build.sh`. It writes the four binaries to `launcher/dist/`.
2. Zip them the way the existing downloads are laid out (Windows: the `.exe` and a README; Linux: `MouseKeyboardTest` and a README; macOS: the `.app` bundle with `mkt-x86_64`, `mkt-arm64` and the small `MouseKeyboardTest` shell script, plus a README). Keep the executable bit on the Linux and macOS files.
3. Put the zips in `downloads/`, then update `SHA256SUMS.txt` and the sizes and checksums shown in the Downloads section of `index.html`. To get a checksum:

```
sha256sum downloads/*.zip
```

Copyright (c) 2026 Shinjuu. All rights reserved. See the terms on the landing page.
