import Application, {IApplication} from './Application';

// This file was autogenerated. Please do not change.
// All changes will be overwrited on commit.
export interface IUpdateAnApplicationResponse {
    values_updated: IApplication;
}

export default class UpdateAnApplicationResponse {
    readonly _values_updated: Application;

    get valuesUpdated(): Application {
        return this._values_updated;
    }

    constructor(props: IUpdateAnApplicationResponse) {
        this._values_updated = new Application(props.values_updated);
    }

    serialize(): IUpdateAnApplicationResponse {
        const data: IUpdateAnApplicationResponse = {
            values_updated: this._values_updated.serialize(),
        };
        return data;
    }

    validate(): string[] {
        const validate = {
            values_updated: this._values_updated.validate().length === 0,
        };
        const isError: string[] = [];
        Object.keys(validate).forEach((key) => {
            if (!(validate as any)[key]) {
                isError.push(key);
            }
        });
        return isError;
    }

    update(
        props: Partial<IUpdateAnApplicationResponse>,
    ): UpdateAnApplicationResponse {
        return new UpdateAnApplicationResponse({...this.serialize(), ...props});
    }
}
