var path = require('path');
var webpack= require('webpack');

module.exports = {
    entry: {
        index: path.resolve(__dirname, 'app/index.jsx'),
    },
    output: {
        path: path.resolve(__dirname, 'build'),
        filename: '[name].js'
    },
    module: {
        loaders: [
            {
                test : /\.jsx$/,
                loader : 'babel-loader',
                query:
                {
                    presets:['es2015', 'react']
                }
            }, {
                test: /\.css$/,
                loader: 'style!css'
            }, {
                test: /\.less$/,
                loader: 'style!css!less'
            }, {
                test: /\.(png|jpg)$/,
                loader: 'file-loader?name=images/[name].[ext]'
            }, {
                test: /\.(woff|woff2|eot|svg|ttf)($|\?)/,
                loader: 'file-loader?name=fonts/[name].[ext]'
            }
        ]
    },
    devtool: 'source-map',
}
