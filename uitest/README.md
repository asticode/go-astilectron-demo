A sample NodeJS / Spectron-based end-to-end test for the Astilectron demo

# Prerequisites

* Install npm (see [https://www.npmjs.com/get-npm]).

# Step 1: install nodejs and the test dependencies

Run the following command:
    $ npm install

# Step 2: run the demo at least once after building it (to allow it to provision the electron artifacts)

# Step 3: run the test:

Run the following command:
    $ npm test

# Troubleshooting

* If you get an error suggesting that the test can't find the electron executable, ensure you run the demo at least once to allow astilectron to provision the electron artifacts

* "Uncaught javascript exception" ECONNRESET: 