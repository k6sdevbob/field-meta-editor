import { createContext } from 'react';
import { makeAutoObservable } from 'mobx';
import { createStoreContextHook, createStoreProviderHook, createStoreSelectorHook, type StoreType } from './utils.js';


class MainStore {

    public darkMode = false;
    
    protected effects: (() => void)[] = [];

    constructor() {
        makeAutoObservable(this, {
            // @ts-expect-error - non-public fields
            effects: false,
        });
    }

    /** Ensure this method to be called in browser */
    public async init() {
        this.darkMode = window.matchMedia('(prefers-color-scheme: dark)').matches;
        this.effects = [];
    }

    public destroy() {
        this.effects.splice(0).forEach(dispose => dispose());
    }

    public setDarkMode(darkMode: boolean) {
        this.darkMode = darkMode;
    }

}

const MainStoreContext = createContext<StoreType<MainStore>>(null!);

export const useMainStoreProvider = createStoreProviderHook<typeof MainStore>(MainStore, MainStoreContext);
export const useMainStore = createStoreContextHook<typeof MainStore>(MainStore, MainStoreContext);
export const useMainStoreSelector = createStoreSelectorHook<typeof MainStore>(MainStore, MainStoreContext);
