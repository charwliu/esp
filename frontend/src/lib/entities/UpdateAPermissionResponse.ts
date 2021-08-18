import Permission, {IPermission} from './Permission';

// This file was autogenerated. Please do not change.
// All changes will be overwrited on commit.
export interface IUpdateAPermissionResponse {
    values_updated: IPermission;
}

export default class UpdateAPermissionResponse {
    readonly _values_updated: Permission;

    get valuesUpdated(): Permission {
        return this._values_updated;
    }

    constructor(props: IUpdateAPermissionResponse) {
        this._values_updated = new Permission(props.values_updated);
    }

    serialize(): IUpdateAPermissionResponse {
        const data: IUpdateAPermissionResponse = {
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
        props: Partial<IUpdateAPermissionResponse>,
    ): UpdateAPermissionResponse {
        return new UpdateAPermissionResponse({...this.serialize(), ...props});
    }
}