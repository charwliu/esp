import AddTrustedApplicationResponse, {
    IAddTrustedApplicationResponse,
} from '@entities/AddTrustedApplicationResponse';
import AssignAPermissionToARoleResponse, {
    IAssignAPermissionToARoleResponse,
} from '@entities/AssignAPermissionToARoleResponse';
import AssignARoleToAOrganizationResponse, {
    IAssignARoleToAOrganizationResponse,
} from '@entities/AssignARoleToAOrganizationResponse';
import AssignARoleToAUserResponse, {
    IAssignARoleToAUserResponse,
} from '@entities/AssignARoleToAUserResponse';
import CreateAnApplicationRequest, {
    ICreateAnApplicationRequest,
} from '@entities/CreateAnApplicationRequest';
import CreateAnApplicationResponse, {
    ICreateAnApplicationResponse,
} from '@entities/CreateAnApplicationResponse';
import CreateAnIotAgentResponse, {
    ICreateAnIotAgentResponse,
} from '@entities/CreateAnIotAgentResponse';
import CreateAPepProxyResponse, {
    ICreateAPepProxyResponse,
} from '@entities/CreateAPepProxyResponse';
import CreateAPermissionRequest, {
    ICreateAPermissionRequest,
} from '@entities/CreateAPermissionRequest';
import CreateAPermissionResponse, {
    ICreateAPermissionResponse,
} from '@entities/CreateAPermissionResponse';
import CreateARoleRequest, {
    ICreateARoleRequest,
} from '@entities/CreateARoleRequest';
import CreateARoleResponse, {
    ICreateARoleResponse,
} from '@entities/CreateARoleResponse';
import InfoPepProxyResponse, {
    IInfoPepProxyResponse,
} from '@entities/InfoPepProxyResponse';
import ListApplicationsResponse, {
    IListApplicationsResponse,
} from '@entities/ListApplicationsResponse';
import ListOfIotAgentsResponse, {
    IListOfIotAgentsResponse,
} from '@entities/ListOfIotAgentsResponse';
import ListOrganizationsInAnApplicationResponse, {
    IListOrganizationsInAnApplicationResponse,
} from '@entities/ListOrganizationsInAnApplicationResponse';
import ListOrganizationsRoleAssignmentsResponse, {
    IListOrganizationsRoleAssignmentsResponse,
} from '@entities/ListOrganizationsRoleAssignmentsResponse';
import ListPermissionsAssociatedToARoleResponse, {
    IListPermissionsAssociatedToARoleResponse,
} from '@entities/ListPermissionsAssociatedToARoleResponse';
import ListPermissionsResponse, {
    IListPermissionsResponse,
} from '@entities/ListPermissionsResponse';
import ListRolesResponse, {
    IListRolesResponse,
} from '@entities/ListRolesResponse';
import ListTrustedApplicationsAssociatedToAnApplicationResponse, {
    IListTrustedApplicationsAssociatedToAnApplicationResponse,
} from '@entities/ListTrustedApplicationsAssociatedToAnApplicationResponse';
import ListUsersInAnApplicationResponse, {
    IListUsersInAnApplicationResponse,
} from '@entities/ListUsersInAnApplicationResponse';
import ListUsersRoleAssignmentsResponse, {
    IListUsersRoleAssignmentsResponse,
} from '@entities/ListUsersRoleAssignmentsResponse';
import ReadApplicationDetailsResponse, {
    IReadApplicationDetailsResponse,
} from '@entities/ReadApplicationDetailsResponse';
import ReadIotAgentDetailsResponse, {
    IReadIotAgentDetailsResponse,
} from '@entities/ReadIotAgentDetailsResponse';
import ReadPermissionDetailsResponse, {
    IReadPermissionDetailsResponse,
} from '@entities/ReadPermissionDetailsResponse';
import ReadRoleDetailsResponse, {
    IReadRoleDetailsResponse,
} from '@entities/ReadRoleDetailsResponse';
import ReadServiceProviderConfigurationResponse, {
    IReadServiceProviderConfigurationResponse,
} from '@entities/ReadServiceProviderConfigurationResponse';
import ResetPasswordOfIotAgentResponse, {
    IResetPasswordOfIotAgentResponse,
} from '@entities/ResetPasswordOfIotAgentResponse';
import ResetPasswordOfPepProxyResponse, {
    IResetPasswordOfPepProxyResponse,
} from '@entities/ResetPasswordOfPepProxyResponse';
import UpdateAnApplicationRequest, {
    IUpdateAnApplicationRequest,
} from '@entities/UpdateAnApplicationRequest';
import UpdateAnApplicationResponse, {
    IUpdateAnApplicationResponse,
} from '@entities/UpdateAnApplicationResponse';
import UpdateAPermissionRequest, {
    IUpdateAPermissionRequest,
} from '@entities/UpdateAPermissionRequest';
import UpdateAPermissionResponse, {
    IUpdateAPermissionResponse,
} from '@entities/UpdateAPermissionResponse';
import UpdateARoleRequest, {
    IUpdateARoleRequest,
} from '@entities/UpdateARoleRequest';
import UpdateARoleResponse, {
    IUpdateARoleResponse,
} from '@entities/UpdateARoleResponse';

// This file was autogenerated. Please do not change.
// All changes will be overwrited on commit.
export default class ApplicationApi {
    static async AddTrustedApplication(): Promise<
        IAddTrustedApplicationResponse | Error
    > {
        return await fetch(
            `/api/v1/applications/:application_id/trusted_applications/:trustedApplicationId`,
            {
                method: 'PUT',
            },
        ).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async AssignOrganizationRole(): Promise<
        IAssignARoleToAOrganizationResponse | Error
    > {
        return await fetch(
            `/api/v1/applications/:application_id/organizations/:organization_id/roles/:role_id/organization_roles/:organization_role_id`,
            {
                method: 'PUT',
            },
        ).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async AssignPermission(): Promise<
        IAssignAPermissionToARoleResponse | Error
    > {
        return await fetch(
            `/api/v1/applications/:application_id/roles/:role_id/permissions/:permission_id`,
            {
                method: 'POST',
            },
        ).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async AssignRole(): Promise<IAssignARoleToAUserResponse | Error> {
        return await fetch(
            `/api/v1/applications/:application_id/users/:user_id/roles/:role_id`,
            {
                method: 'PUT',
            },
        ).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async CreateApplication(
        createanapplicationrequest: ICreateAnApplicationRequest,
    ): Promise<ICreateAnApplicationResponse | string[] | Error> {
        const haveError: string[] = [];
        const createanapplicationrequestValid = new CreateAnApplicationRequest(
            createanapplicationrequest,
        );
        haveError.push(...createanapplicationrequestValid.validate());
        if (haveError.length > 0) {
            return Promise.resolve(haveError);
        }
        return await fetch(`/api/v1/applications`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(createanapplicationrequestValid.serialize()),
        }).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async CreateIotAgent(): Promise<ICreateAnIotAgentResponse | Error> {
        return await fetch(`/api/v1/applications/:application_id/iot_agents`, {
            method: 'POST',
        }).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async CreatePepProxy(): Promise<ICreateAPepProxyResponse | Error> {
        return await fetch(`/api/v1/applications/:application_id/pep_proxies`, {
            method: 'POST',
        }).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async CreatePermission(
        createapermissionrequest: ICreateAPermissionRequest,
    ): Promise<ICreateAPermissionResponse | string[] | Error> {
        const haveError: string[] = [];
        const createapermissionrequestValid = new CreateAPermissionRequest(
            createapermissionrequest,
        );
        haveError.push(...createapermissionrequestValid.validate());
        if (haveError.length > 0) {
            return Promise.resolve(haveError);
        }
        return await fetch(`/api/v1/applications/:application_id/permissions`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(createapermissionrequestValid.serialize()),
        }).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async CreateRole(
        createarolerequest: ICreateARoleRequest,
    ): Promise<ICreateARoleResponse | string[] | Error> {
        const haveError: string[] = [];
        const createarolerequestValid = new CreateARoleRequest(
            createarolerequest,
        );
        haveError.push(...createarolerequestValid.validate());
        if (haveError.length > 0) {
            return Promise.resolve(haveError);
        }
        return await fetch(`/api/v1/applications/:application_id/roles`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(createarolerequestValid.serialize()),
        }).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async DeleteApplication(): Promise<number | Error> {
        return await fetch(`/api/v1/applications/:application_id`, {
            method: 'DELETE',
        }).then(async (res) => {
            if (res.status === 200) {
                return res.status;
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async DeleteIotAgent(): Promise<number | Error> {
        return await fetch(
            `/api/v1/applications/:application_id/permissions/:iot_agent_id`,
            {
                method: 'DELETE',
            },
        ).then(async (res) => {
            if (res.status === 200) {
                return res.status;
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async DeletePepProxy(): Promise<number | Error> {
        return await fetch(`/api/v1/applications/:application_id/pep_proxies`, {
            method: 'DELETE',
        }).then(async (res) => {
            if (res.status === 200) {
                return res.status;
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async DeletePermission(): Promise<number | Error> {
        return await fetch(
            `/api/v1/applications/:application_id/permissions/:permission_id`,
            {
                method: 'DELETE',
            },
        ).then(async (res) => {
            if (res.status === 200) {
                return res.status;
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async DeleteRole(): Promise<number | Error> {
        return await fetch(
            `/api/v1/applications/:application_id/roles/:role_id`,
            {
                method: 'DELETE',
            },
        ).then(async (res) => {
            if (res.status === 200) {
                return res.status;
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async InfoPepProxy(): Promise<IInfoPepProxyResponse | Error> {
        return await fetch(`/api/v1/applications/:application_id/pep_proxies`, {
            method: 'GET',
        }).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async ListApplicationOrganizations(): Promise<
        IListOrganizationsInAnApplicationResponse | Error
    > {
        return await fetch(
            `/api/v1/applications/:application_id/organizations`,
            {
                method: 'GET',
            },
        ).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async ListApplications(): Promise<
        IListApplicationsResponse | Error
    > {
        return await fetch(`/api/v1/applications`, {
            method: 'GET',
        }).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async ListIotAgents(): Promise<IListOfIotAgentsResponse | Error> {
        return await fetch(`/api/v1/applications/:application_id/iot_agents`, {
            method: 'GET',
        }).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async ListPermissions(): Promise<IListPermissionsResponse | Error> {
        return await fetch(`/api/v1/applications/:application_id/permissions`, {
            method: 'GET',
        }).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async ListRoleAssignments(): Promise<
        IListOrganizationsRoleAssignmentsResponse | Error
    > {
        return await fetch(
            `/api/v1/applications/:application_id/organizations/:organization_id/roles`,
            {
                method: 'GET',
            },
        ).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async ListRolePermissions(): Promise<
        IListPermissionsAssociatedToARoleResponse | Error
    > {
        return await fetch(
            `/api/v1/applications/:application_id/roles/:role_id/permissions`,
            {
                method: 'GET',
            },
        ).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async ListRoles(): Promise<IListRolesResponse | Error> {
        return await fetch(`/api/v1/applications/:application_id/roles`, {
            method: 'GET',
        }).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async ListTrustedApplications(): Promise<
        IListTrustedApplicationsAssociatedToAnApplicationResponse | Error
    > {
        return await fetch(
            `/api/v1/applications/:application_id/trusted_applications`,
            {
                method: 'GET',
            },
        ).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async ListUsersInApplication(): Promise<
        IListUsersInAnApplicationResponse | Error
    > {
        return await fetch(`/api/v1/applications/:application_id/users`, {
            method: 'GET',
        }).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async ListUsersRoleAssignments(): Promise<
        IListUsersRoleAssignmentsResponse | Error
    > {
        return await fetch(
            `/api/v1/applications/:application_id/users/:user_id/roles`,
            {
                method: 'GET',
            },
        ).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async ReadApplication(): Promise<
        IReadApplicationDetailsResponse | Error
    > {
        return await fetch(`/api/v1/applications/:application_id`, {
            method: 'GET',
        }).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async ReadConfiguration(): Promise<
        IReadServiceProviderConfigurationResponse | Error
    > {
        return await fetch(`/api/v1/service_providers/configs`, {
            method: 'GET',
        }).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async ReadIotAgent(): Promise<IReadIotAgentDetailsResponse | Error> {
        return await fetch(
            `/api/v1/applications/:application_id/permissions/:iot_agent_id`,
            {
                method: 'GET',
            },
        ).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async ReadPermission(): Promise<
        IReadPermissionDetailsResponse | Error
    > {
        return await fetch(
            `/api/v1/applications/:application_id/permissions/:permission_id`,
            {
                method: 'GET',
            },
        ).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async ReadRoleDetails(): Promise<IReadRoleDetailsResponse | Error> {
        return await fetch(
            `/api/v1/applications/:application_id/roles/:role_id`,
            {
                method: 'GET',
            },
        ).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async RemovePermission(): Promise<number | Error> {
        return await fetch(
            `/api/v1/applications/:application_id/roles/:role_id/permissions/:permission_id`,
            {
                method: 'DELETE',
            },
        ).then(async (res) => {
            if (res.status === 200) {
                return res.status;
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async RemoveRoleAssignment(): Promise<number | Error> {
        return await fetch(
            `/api/v1/applications/:application_id/users/:user_id/roles/:role_id`,
            {
                method: 'DELETE',
            },
        ).then(async (res) => {
            if (res.status === 200) {
                return res.status;
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async RemoveRoleAssignmentFromOrgization(): Promise<number | Error> {
        return await fetch(
            `/api/v1/applications/:application_id/organizations/:organization_id/roles/:role_id/organization_roles/:organization_role_id`,
            {
                method: 'DELETE',
            },
        ).then(async (res) => {
            if (res.status === 200) {
                return res.status;
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async RemoveTrustedApplication(): Promise<number | Error> {
        return await fetch(
            `/api/v1/applications/:application_id/trusted_applications/:trustedApplicationId`,
            {
                method: 'DELETE',
            },
        ).then(async (res) => {
            if (res.status === 200) {
                return res.status;
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async ResetPassword(): Promise<
        IResetPasswordOfPepProxyResponse | Error
    > {
        return await fetch(`/api/v1/applications/:application_id/pep_proxies`, {
            method: 'PATCH',
        }).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async ResetPasswordOfIotAgent(): Promise<
        IResetPasswordOfIotAgentResponse | Error
    > {
        return await fetch(
            `/api/v1/applications/:application_id/permissions/:iot_agent_id`,
            {
                method: 'PATCH',
            },
        ).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async UpdateApplication(
        updateanapplicationrequest: IUpdateAnApplicationRequest,
    ): Promise<IUpdateAnApplicationResponse | string[] | Error> {
        const haveError: string[] = [];
        const updateanapplicationrequestValid = new UpdateAnApplicationRequest(
            updateanapplicationrequest,
        );
        haveError.push(...updateanapplicationrequestValid.validate());
        if (haveError.length > 0) {
            return Promise.resolve(haveError);
        }
        return await fetch(`/api/v1/applications/:application_id`, {
            method: 'PATCH',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(updateanapplicationrequestValid.serialize()),
        }).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async UpdatePermission(
        updateapermissionrequest: IUpdateAPermissionRequest,
    ): Promise<IUpdateAPermissionResponse | string[] | Error> {
        const haveError: string[] = [];
        const updateapermissionrequestValid = new UpdateAPermissionRequest(
            updateapermissionrequest,
        );
        haveError.push(...updateapermissionrequestValid.validate());
        if (haveError.length > 0) {
            return Promise.resolve(haveError);
        }
        return await fetch(
            `/api/v1/applications/:application_id/permissions/:permission_id`,
            {
                method: 'PATCH',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(updateapermissionrequestValid.serialize()),
            },
        ).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }

    static async UpdateRole(
        updatearolerequest: IUpdateARoleRequest,
    ): Promise<IUpdateARoleResponse | string[] | Error> {
        const haveError: string[] = [];
        const updatearolerequestValid = new UpdateARoleRequest(
            updatearolerequest,
        );
        haveError.push(...updatearolerequestValid.validate());
        if (haveError.length > 0) {
            return Promise.resolve(haveError);
        }
        return await fetch(
            `/api/v1/applications/:application_id/roles/:role_id`,
            {
                method: 'PATCH',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(updatearolerequestValid.serialize()),
            },
        ).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        });
    }
}
