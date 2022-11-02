const child_process = require("child_process");
const process = require("process");
const fs = require("fs");
const path = require("path");

const TESTER_RELEASE_URL =
  "https://github.com/massalabs/massa-sc-tester/releases/download";
const VERSION = "v0.3.1";

console.log(`Downloading simulator binary for ${process.platform}`);
let fileName;
switch (process.platform) {
  case "darwin":
    fileName = "release_macos.tar.gz";
    break;
  case "linux":
    fileName = "release_linux.tar.gz";
    break;
  case "win32":
  default:
    console.error(`OS not supported`);
    process.exit(1);
}

const url = `${TESTER_RELEASE_URL}/${VERSION}/${fileName}`;

let outputDir = path.join(process.cwd(), "simulator");
fs.mkdirSync(outputDir, { recursive: true });

const cmd = `wget -qO- ${url} | tar xz -C ${outputDir} --strip-components 1`;
child_process.execSync(cmd);

process.exit(0);
