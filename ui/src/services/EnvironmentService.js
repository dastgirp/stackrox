import axios from 'axios';
import queryString from 'query-string';

const baseUrl = '/v1/networkgraph';
const networkPoliciesUrl = '/v1/networkpolicies';

/**
 * Fetches nodes and links for the environment graph.
 * Returns response with nodes and links
 *
 * @returns {Promise<Object, Error>}
 */
export function fetchEnvironmentGraph(filters) {
    const params = queryString.stringify({
        ...filters
    });
    return axios.get(`${baseUrl}?${params}`).then(response => ({
        response: response.data
    }));
}

/**
 * Fetches policies details for given array of ids.
 *
 * @param {!array} policyIds
 * @returns {Promise<Object, Error>}
 */
export function fetchNetworkPolicies(policyIds) {
    const networkPoliciesPromises = policyIds.map(policyId =>
        axios.get(`${networkPoliciesUrl}/${policyId}`)
    );
    return axios
        .all([...networkPoliciesPromises])
        .then(response => ({ response: response.map(networkPolicy => networkPolicy.data) }));
}
