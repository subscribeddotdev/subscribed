module.exports = {
  plugins: {
    "postcss-normalize": {},
    "@csstools/postcss-global-data": {
      files: ["src/common/styles/mq.css"],
    },
    "postcss-preset-env": {
      stage: 0,
    },
  },
};
