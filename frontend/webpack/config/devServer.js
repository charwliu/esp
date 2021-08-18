/**
 * Created by: Andrey Polyakov (andrey@polyakov.im)
 * @see https://webpack.js.org/configuration/dev-server/
 */
import isWindows from 'is-windows';

import {devServerProxyConfig} from './devServierProxy';

const defaultPort = 8080;

const devServerHost = isWindows() ? '127.0.0.1' : '0.0.0.0';

export const devServerUrl = `http://${devServerHost}:${defaultPort}/`;

export const devServerConfig = {
    port: defaultPort,
    historyApiFallback: true,
    headers: {'Access-Control-Allow-Origin': '*'},
    proxy: devServerProxyConfig,
    hot: true,
    client: {
        overlay: false,
    },
    host: devServerHost,
    devMiddleware: {
        publicPath: '/',
    },
};
