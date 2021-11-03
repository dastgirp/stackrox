import withAuth from '../../helpers/basicAuth';
import { url, selectors } from '../../constants/VulnManagementPage';
import {
    hasExpectedHeaderColumns,
    allChecksForEntities,
    allCVECheck,
    allFixableCheck,
} from '../../helpers/vmWorkflowUtils';
import { hasFeatureFlag } from '../../helpers/features';
import * as api from '../../constants/apiEndpoints';

describe('Images list page and its entity detail page, related entities sub list validations ', () => {
    withAuth();

    const vulnRiskCveTabsEnabled = hasFeatureFlag('ROX_VULN_RISK_MANAGEMENT');

    it('should display all the columns and links expected in images list page', () => {
        cy.intercept('POST', api.vulnMgmt.graphqlEntities('IMAGE')).as('getImages');
        cy.visit(url.list.images);
        cy.wait('@getImages');

        hasExpectedHeaderColumns([
            'Image',
            'CVEs',
            'Top CVSS',
            'Created',
            'Scan Time',
            'Image Status',
            'Deployments',
            'Components',
            'Risk Priority',
        ]);
        cy.get(selectors.tableBodyColumn).each(($el) => {
            const columnValue = $el.text().toLowerCase();
            if (columnValue !== 'no deployments' && columnValue.includes('Deployment')) {
                allChecksForEntities(url.list.images, 'deployment');
            }
            if (columnValue !== 'no components' && columnValue.includes('Component')) {
                allChecksForEntities(url.list.images, 'component');
            }
            if (columnValue !== 'no cves' && columnValue.includes('fixable')) {
                if (vulnRiskCveTabsEnabled) {
                    cy.get(`${selectors.tableBodyColumn}:eq(0)`).click({ force: true });

                    cy.get('.pf-c-tabs .pf-c-tabs__item:eq(0):contains("Observed CVEs")').click();

                    /* 

                    @TODO: Uncomment when the tables are made
                    
                    cy.get('.pf-c-tabs .pf-c-tabs__item:eq(1):contains("Deferred CVEs")').click();

                    cy.get('.pf-c-tabs .pf-c-tabs__item:eq(2):contains("False Positive CVEs")').click(); 
                        
                    */
                } else {
                    allFixableCheck(url.list.images);
                }
            }
            if (columnValue !== 'no cves' && columnValue.includes('cve')) {
                allCVECheck(url.list.images);
            }
        });
        //  TBD to be fixed after back end sorting is fixed
        //  validateSort(selectors.riskScoreCol);
    });

    it('should show entity icon, not back button, if there is only one item on the side panel stack', () => {
        cy.intercept('POST', api.vulnMgmt.graphqlEntities('IMAGE')).as('getImages');
        cy.visit(url.list.images);
        cy.wait('@getImages');

        cy.get(`${selectors.deploymentCountLink}:eq(0)`).click({ force: true });
        cy.wait(1000);
        cy.get(selectors.backButton).should('exist');
        cy.get(selectors.entityIcon).should('not.exist');

        cy.get(selectors.backButton).click();
        cy.get(selectors.backButton).should('not.exist');
        cy.get(selectors.entityIcon).should('exist');
    });
});
