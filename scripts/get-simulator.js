const child_process = require('child_process');
const process = require('process');

const TESTER_RELEASE_URL =
  "https://github.com/massalabs/massa-sc-tester/releases/download";
const VERSION = "v0.3.1";

let fileName;
switch (process.platform) {
  case "win32":
    fileName = "release_windows.zip";
    process.exit(0);
  case "darwin":
    fileName = "release_macos.tar.gz";
    break;
  case "linux":
    fileName = "release_linux.tar.gz";
    break;
  default:
    console.error(`OS not supported`);
    process.exit(1);
}

const url = `${TESTER_RELEASE_URL}/${VERSION}/${fileName}`;

let outputDir = "./"
if (process.argv[2] && process.argv[2] === "-o" ) {
  outputDir += `/${process.argv[3]}`
}

let cmd;
if (process.platform === "win32") {
    cmd = `powershell -Command "Invoke-WebRequest -Uri ${url} -OutFile ${fileName}"| Expand-Archive -Path ${fileName} -DestinationPath ${outputDir}/simulator`;
} else {
    cmd = `wget -qO- ${url} | tar xz -C ${outputDir}/simulator --strip-components 1`;
}

child_process.execSync(cmd);

process.exit(0);
