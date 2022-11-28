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
child_process.execSync(`curl -Ls ${url} -o ${fileName}`);

let cmd = `tar -xzf ${fileName} -C ${outputDir} --strip-components 1`;
if (process.platform === "win32") {
  cmd += ` && del ${fileName}`;
} else {
  if (platformName === "Windows") {
    // special case for the CI cross compile
    cmd = `unzip -j ${fileName}   "massa/massa-sc-tester.exe" -d ${outputDir}`;
  }
  cmd += ` && rm ${fileName}`;
}
console.log(`Extracting simulator binary in ${outputDir}`);
child_process.execSync(cmd);

process.exit(0);
