// Semicolon on separate line following the strings prevents an extra changed line to add a string at the end.
// However, add strings in alphabetical order to minimize merge conflicts when multiple people add strings.
// prettier-ignore
export type FeatureFlagEnvVar =
    | 'ROX_DECOMMISSIONED_CLUSTER_RETENTION'
    | 'ROX_NETPOL_FIELDS'
    | 'ROX_NEW_POLICY_CATEGORIES'
    | 'ROX_QUAY_ROBOT_ACCOUNTS'
    | 'ROX_SEARCH_PAGE_UI'
    | 'ROX_SYSTEM_HEALTH_PF'
    | 'ROX_FRONTEND_VM_UPDATES'
    ;
