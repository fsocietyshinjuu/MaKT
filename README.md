Mouse and Keyboard Test (desktop app)

1. Install Node.js (https://nodejs.org), version 18 or newer.
2. Open a terminal in this folder and run:
     npm install
3. Run it right away:
     npm start
   Or build a standalone app (output goes to the "dist" folder):
     npm run build:win      (Windows: dist\MouseKeyboardTest-win32-x64\MouseKeyboardTest.exe)
     npm run build:mac      (macOS: dist/MouseKeyboardTest-darwin-*/MouseKeyboardTest.app)
     npm run build:linux    (Linux: dist/MouseKeyboardTest-linux-x64/MouseKeyboardTest)

Copy or zip the whole output folder to move it to another computer. The exe needs the files next to it.
On macOS the app is unsigned: right-click it, choose Open, then confirm.
The page itself is index.html, so edit that file to change the tester.
