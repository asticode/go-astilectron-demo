const Application = require('spectron').Application;
const chai = require('chai');
const chaiAsPromised = require('chai-as-promised');
const electron = require('electron');
const { exec } = require("child_process");

const APPNAME = 'Astilectron demo';
const PORT = 55555; // the port the main process will listen to

global.before(() => {
  chai.should();
  chai.use(chaiAsPromised);
});

// Map nodejs arch to golang arch
let archMap = {
  "arm": "arm",
  "ia32": "386",
  "x86": "386",
  "x64": "amd64",
  "ia64": "amd64"
};

if (archMap[process.arch] === undefined) {
  console.log(`FATAL: unhandled platform/processor type (${process.arch}) - add your variant to archMap in test/hooks.js`);
  process.exit(1);
}

function mainExe() {
  if (process.platform === 'darwin') {
    return `../output/darwin-${archMap[process.arch]}/${APPNAME}.app/Contents/MacOS/${APPNAME}`;
  } else if (process.platform === 'linux') {
    return `../output/linux-${archMap[process.arch]}/${APPNAME}`;
  } else if (process.platform === 'win32') {
    return `../output/windows-${archMap[process.arch]}/${APPNAME}.exe`;
  } else {
    console.log("FATAL: unhandled platform/os - add your variant here");
    process.exit(1);
  }
}

function electronExe() {
  if (process.platform === 'darwin') {
    return `../output/darwin-${archMap[process.arch]}/${APPNAME}.app/Contents/MacOS/vendor/electron-darwin-${archMap[process.arch]}/${APPNAME}.app/Contents/MacOS/${APPNAME}`;
  } else if (process.platform === 'linux') {
    return `../output/linux-${archMap[process.arch]}/vendor/electron-linux-${archMap[process.arch]}/electron`;
  } else if (process.platform === 'win32') {
    return `${process.env.APPDATA}/${APPNAME}/vendor/electron-windows-${archMap[process.arch]}/Electron.exe`;
  } else {
    console.log("FATAL: unhandled platform - add your variant here");
    process.exit(1);
  }
}

function astilectronJS() {
  if (process.platform === 'darwin') {
    return `../output/darwin-${archMap[process.arch]}/${APPNAME}.app/Contents/MacOS/vendor/astilectron/main.js`;
  } else if (process.platform === 'linux') {
    return `../output/linux-${archMap[process.arch]}/vendor/vendor/astilectron/main.js`;
  } else if (process.platform === 'win32') {
    return `${process.env.APPDATA}/${APPNAME}/vendor/astilectron/main.js`;
  } else {
    console.log("FATAL: unhandled platform - add your variant here");
    process.exit(1);
  }
}

module.exports = {
  async startMainApp() {
    console.log(`node arch: "${process.arch}"   golang arch: "${archMap[process.arch]}"`)
    console.log(`Starting main exe: ${mainExe()}`);
    exec(`"${mainExe()}" -UITEST ${PORT}`, (error, stdout, stderr) => {
      if (error) {
        console.log(`error: ${error.message}`);
        return;
      }
      if (stderr) {
        console.log(`stderr: ${stderr}`);
        return;
      }
      console.log(`stdout: ${stdout}`);

    });
  },

  async getApp() {
    return module.exports.app;
  },

  async startApp() {
    module.exports.startMainApp();

    console.log(`Starting electron exe: ${electronExe()}`);
    const rendererApp = await new Application({

      path: electronExe(),
      args: [astilectronJS(), `127.0.0.1:${PORT}`, 'true'],

      // for debugging:
      //chromeDriverLogPath: './chromedriver.log',
      //webdriverLogPath: './webdriver.log'

    }).start();
    chaiAsPromised.transferPromiseness = rendererApp.transferPromiseness;
    module.exports.app = rendererApp;
    return rendererApp;
  },

  async stopApp(app) {
    if (app && app.isRunning()) {
      await app.stop();
    }
  }
};
