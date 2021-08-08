import { flow, makeAutoObservable, observable } from 'mobx';
import { Store } from 'Store';

import statsApi from 'Apis/stats';
import queryApi from 'Apis/log';
import safeBrowsingApi from 'Apis/safebrowsing';
import filteringApi from 'Apis/filtering';
import parentalApi from 'Apis/parental';
import safesearchApi from 'Apis/safesearch';

import StatsConfig, { IStatsConfig } from 'Entities/StatsConfig';
import QueryLogConfig, { IQueryLogConfig } from 'Entities/QueryLogConfig';
import FilterConfig, { IFilterConfig } from 'Entities/FilterConfig';
import FilterStatus, { IFilterStatus } from 'Entities/FilterStatus';

import { errorChecker } from 'Helpers/apiErrors';

import { IStore } from './utils';

export default class SomeStore implements IStore {
    rootStore: Store;

    inited = false;

    statsConfig: StatsConfig | undefined;

    queryLogConfig: QueryLogConfig | undefined;

    safebrowsing: boolean | undefined;

    filteringConfig: FilterConfig | undefined;

    parental: boolean | undefined;

    safesearch: boolean | undefined;

    constructor(rootStore: Store) {
        this.rootStore = rootStore;
        makeAutoObservable(this, {
            rootStore: false,
            inited: observable,
            init: flow,

            statsConfig: observable.ref,
            queryLogConfig: observable.ref,
            safebrowsing: observable,
            filteringConfig: observable.ref,
            parental: observable,
            safesearch: observable,

            updateStatsConfig: flow,
            statsInfo: flow,
            statsReset: flow,
            updateQueryLogConfig: flow,
            queryLogInfo: flow,
            querylogClear: flow,
            safebrowsingDisable: flow,
            safebrowsingEnable: flow,
            safebrowsingStatus: flow,
            updateFilteringConfig: flow,
            filteringStatus: flow,
            parentalDisable: flow,
            parentalEnable: flow,
            parentalStatus: flow,
            safesearchDisable: flow,
            safesearchEnable: flow,
            safesearchStatus: flow,
        });
    }

    * init() {
        yield this.statsInfo();
        yield this.queryLogInfo();
        yield this.safebrowsingStatus();
        yield this.filteringStatus();
        yield this.parentalStatus();
        yield this.safesearchStatus();
        this.inited = yield true;
    }

    * updateStatsConfig(statsconfig: IStatsConfig) {
        const response: number | Error | string[] = yield statsApi.statsConfig(statsconfig);
        const { result } = errorChecker<number>(response);
        if (result) {
            yield this.statsInfo();
        }
    }

    * statsInfo() {
        const response: IStatsConfig | Error = yield statsApi.statsInfo();
        const { result } = errorChecker<IStatsConfig>(response);
        if (result) {
            this.statsConfig = new StatsConfig(result);
        }
    }

    * statsReset() {
        const response: number | Error = yield statsApi.statsReset();
        const { result } = errorChecker(response);
        if (result) {
            yield this.statsInfo();
            return true;
        }
    }

    * updateQueryLogConfig(querylogconfig: IQueryLogConfig) {
        const response: Error | Number | string[] = yield queryApi.queryLogConfig(querylogconfig);
        const { result } = errorChecker<number>(response);
        if (result) {
            yield this.queryLogInfo();
        }
    }

    * queryLogInfo() {
        const response: Error | IQueryLogConfig = yield queryApi.queryLogInfo();
        const { result } = errorChecker<IQueryLogConfig>(response);
        if (result) {
            this.queryLogConfig = new QueryLogConfig(result);
        }
    }

    * querylogClear() {
        const response: number | Error = yield queryApi.querylogClear();
        const { result } = errorChecker(response);
        if (result) {
            yield this.queryLogInfo();
        }
    }

    * safebrowsingDisable() {
        const response: number | Error = yield safeBrowsingApi.safebrowsingDisable();
        const { result } = errorChecker<number>(response);
        if (result) {
            this.safebrowsing = false;
        }
    }

    * safebrowsingEnable() {
        const response: number | Error = yield safeBrowsingApi.safebrowsingEnable();
        const { result } = errorChecker(response);
        if (result) {
            this.safebrowsing = true;
        }
    }

    * safebrowsingStatus() {
        const response: { enabled: boolean } | Error = yield safeBrowsingApi.safebrowsingStatus();
        const { result } = errorChecker(response);
        if (result) {
            this.safebrowsing = result.enabled;
        }
    }

    * updateFilteringConfig(filterconfig: IFilterConfig) {
        const response: number | Error = yield filteringApi.filteringConfig(filterconfig);
        const { result } = errorChecker<number>(response);
        if (result) {
            yield this.filteringStatus();
        }
    }

    * filteringStatus() {
        const response: Error | IFilterStatus = yield filteringApi.filteringStatus();
        const { result } = errorChecker<IFilterStatus>(response);
        if (result) {
            this.filteringConfig = new FilterStatus(result);
        }
    }

    * parentalDisable() {
        const response: number | Error = yield parentalApi.parentalDisable();
        const { result } = errorChecker(response);
        if (result) {
            this.parental = false;
        }
    }

    * parentalEnable() {
        // TODO: remove magic;
        const response: number | Error = yield parentalApi.parentalEnable('sensitivity=TEEN');
        const { result } = errorChecker(response);
        if (result) {
            this.parental = true;
        }
    }

    * parentalStatus() {
        const response: { enabled: boolean } | Error = yield parentalApi.parentalStatus();
        const { result } = errorChecker(response);
        if (result) {
            this.parental = result.enabled;
        }
    }

    * safesearchDisable() {
        const response: number | Error = yield safesearchApi.safesearchDisable();
        const { result } = errorChecker(response);
        if (result) {
            this.safesearch = false;
        }
    }

    * safesearchEnable() {
        const response: number | Error = yield safesearchApi.safesearchEnable();
        const { result } = errorChecker(response);
        if (result) {
            this.safesearch = true;
        }
    }

    * safesearchStatus() {
        const response: { enabled: boolean } | Error = yield safesearchApi.safesearchStatus();
        const { result } = errorChecker(response);
        if (result) {
            this.safesearch = result.enabled;
        }
    }
}
