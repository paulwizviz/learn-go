const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const CopyPlugin = require('copy-webpack-plugin');

const BUILD_DIR = path.resolve(__dirname, '..', 'public');
const SRC_DIR = path.resolve(__dirname, '..', 'src');

module.exports = {
  entry: `${SRC_DIR}/index.js`,
  output: {
    path: BUILD_DIR,
    filename: 'bundle.js',
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: 'babel-loader',
      },
    ],
  },
  plugins: [
    new HtmlWebpackPlugin({ 
      template: path.resolve(SRC_DIR,'index.html'),
      inject: 'body'
    }),
    new CopyPlugin([
      { from: path.resolve(__dirname,'..','images'), to: path.resolve(BUILD_DIR,'images') },
    ]),
  ]
}