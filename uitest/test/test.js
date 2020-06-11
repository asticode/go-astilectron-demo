const hooks = require('./hooks');
const config = require('../config').get(process.env.NODE_ENV);

var app;

describe('Setup', () => {

  before(async () => {
    app = await hooks.startApp();
  });

  after(async () => {
    await hooks.stopApp(app);
  });

  it('opens a window', async () => {
    await app.client
      .waitUntilWindowLoaded()
      .getWindowCount()
      .should.eventually.be.above(0)

      .getTitle().should.eventually.equal('') // the demo doesnt set a title
  });

  it('finds some files', async () => {
    await app.client
      .getText('#files_count').should.eventually.not.be.equal('')
     });

});
