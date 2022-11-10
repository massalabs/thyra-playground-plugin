const child_process = require("child_process");
const process = require("process");
const fs = require("fs");
const path = require("path");

const TESTER_RELEASE_URL =
  "https://github.com/massalabs/massa-sc-tester/releases/download";
const VERSION = "v0.3.2";

// env var used by the CI
const GOOS = process.env["GOOS"];
const GOARCH = process.env["GOARCH"];

let fileName;
let platformName;
if (process.platform === "win32" || GOOS === "windows") {
  fileName = "release_windows.zip";
  platformName = "Windows";
} else if (process.platform === "darwin" || GOOS === "darwin") {
  if (process.arch === "arm64" || GOARCH === "arm64") {
    fileName = "release_macos-arm64.tar.gz";
    platformName = "MacOs arm64";
  } else {
    fileName = "release_macos-amd64.tar.gz";
    platformName = "MacOs amd64";
  }
} else {
  platformName = "Linux";
  fileName = "release_linux.tar.gz";
}

const url = `${TESTER_RELEASE_URL}/${VERSION}/${fileName}`;

let outputDir = path.join(process.cwd(), "simulator");
fs.mkdirSync(outputDir, { recursive: true });

console.log(`Downloading simulator binary for ${platformName} at url: ${url}`);
child_process.execSync(`curl -Lo ${fileName} ${url}`);

let cmd;
if (platformName === "Windows") {
  cmd = `unzip -j ${fileName}  "massa/massa-sc-tester.exe" -d ${outputDir}`;
} else {
  cmd = `tar xf ${fileName} -C ${outputDir} --strip-components 1`;
}
console.log(`Extracting simulator binary in ${outputDir}`);
child_process.execSync(cmd);

process.exit(0);
