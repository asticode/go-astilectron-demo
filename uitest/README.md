A sample NodeJS / Spectron-based end-to-end test for the Astilectron demo

# Prerequisites

* Install npm (see (https://www.npmjs.com/get-npm)).

# Step 1: install nodejs and the test dependencies

Run the following command:
    $ npm install

# Step 2: run the demo at least once after building it (to allow it to provision the electron artifacts)

# Step 3: run the test:

Run the following command:
    $ npm test

# Troubleshooting

* If you get an error suggesting that the test can't find the electron executable, ensure you run the demo at least once to allow astilectron to provision the electron artifacts

* If the "before all" hook times out, the test is having trouble starting the chromedriver to connect to electron. (you may see an error message saying DevToolsActivePort does not exist). A common problem is a mis-match of the version of electron Astilectron is using vs. the version of the chromedriver bundled with spectron.  See the version compatibility matrix at (https://github.com/electron-userland/spectron) and adjust the spectron version accordingly in package.json and rerun "npm install".  

* If all else fails, uncomment chromedriver and webdriver log settings test/hooks.js, rerun the test and then scour the logs.
