import OrganizationUser, {IOrganizationUser} from './OrganizationUser';

// This file was autogenerated. Please do not change.
// All changes will be overwrited on commit.
export interface IListMembersResponse {
    organization_users: IOrganizationUser[];
}

export default class ListMembersResponse {
    readonly _organization_users: OrganizationUser[];

    get organizationUsers(): OrganizationUser[] {
        return this._organization_users;
    }

    constructor(props: IListMembersResponse) {
        this._organization_users = props.organization_users.map(
            (p) => new OrganizationUser(p),
        );
    }

    serialize(): IListMembersResponse {
        const data: IListMembersResponse = {
            organization_users: this._organization_users.map((p) =>
                p.serialize(),
            ),
        };
        return data;
    }

    validate(): string[] {
        const validate = {
            organization_users: this._organization_users.reduce(
                (result, p) => result && p.validate().length === 0,
                true,
            ),
        };
        const isError: string[] = [];
        Object.keys(validate).forEach((key) => {
            if (!(validate as any)[key]) {
                isError.push(key);
            }
        });
        return isError;
    }

    update(props: Partial<IListMembersResponse>): ListMembersResponse {
        return new ListMembersResponse({...this.serialize(), ...props});
    }
}