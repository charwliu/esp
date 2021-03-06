import Organization, {IOrganization} from './Organization';

// This file was autogenerated. Please do not change.
// All changes will be overwrited on commit.
export interface IReadInfoAboutAnOrganizationResponse {
    organization: IOrganization;
}

export default class ReadInfoAboutAnOrganizationResponse {
    readonly _organization: Organization;

    get organization(): Organization {
        return this._organization;
    }

    constructor(props: IReadInfoAboutAnOrganizationResponse) {
        this._organization = new Organization(props.organization);
    }

    serialize(): IReadInfoAboutAnOrganizationResponse {
        const data: IReadInfoAboutAnOrganizationResponse = {
            organization: this._organization.serialize(),
        };
        return data;
    }

    validate(): string[] {
        const validate = {
            organization: this._organization.validate().length === 0,
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
        props: Partial<IReadInfoAboutAnOrganizationResponse>,
    ): ReadInfoAboutAnOrganizationResponse {
        return new ReadInfoAboutAnOrganizationResponse({
            ...this.serialize(),
            ...props,
        });
    }
}
