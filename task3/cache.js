var redisService = require('lib/redisService');

//We should initialize constants while defining. If we get it from configuration file or somewhere else, we can use 'var' or 'let' keyword.
// const cache = redisService.init(port:6379 host:redis.domain.com) without init or
// let cache;

const cache;

function init(opts) {

  cache = redisService.init({
    port: opts.port,
    host: opts.host
  });

}

function get(key) {
  return cache.get(key);
}

function set(key, value, ttl) {
  return cache.set(key, value, ttl || opts.stdTTL);
}

function del(key) {
  return cache.del(key);
}

function flush() {
  return cache.flush();
}

function getCache(opts) {
  return {
    get: get,
    set: set,
    del: del,
    flush: flush
  };
}

module.exports = {
  init: init,
  getCache: getCache
};