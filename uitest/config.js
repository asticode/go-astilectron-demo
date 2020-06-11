const config = {
  default: {
    url: 'https://duckduckgo.com'
  },
};

exports.get = function get(env) {
  return config[env] || config.default;
};