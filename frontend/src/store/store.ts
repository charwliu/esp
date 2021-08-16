import {createContext} from 'react';

export class Store {
    constructor() {}

    init() {}
}

export const storeValue = new Store();

const StoreContext = createContext<Store>(storeValue);

export default StoreContext;
