import { DataSourceInstanceSettings, ScopedVars } from '@grafana/data';
import { DataSourceWithBackend, getTemplateSrv } from '@grafana/runtime';

import { NOAAQuery, NOAAOptions, TCProduct } from './types';

export class DataSource extends DataSourceWithBackend<NOAAQuery, NOAAOptions> {
  readonly options: NOAAOptions;

  constructor(instanceSettings: DataSourceInstanceSettings<NOAAOptions>) {
    super(instanceSettings);
    this.options = instanceSettings.jsonData;
  }

  // This will support annotation queries for 7.2+
  annotations = {};

  /**
   * Do not execute queries that do not exist yet
   */
  filterQuery(query: NOAAQuery): boolean {
    return !!query.product && !!query.station;
  }

  getQueryDisplayText(query: NOAAQuery): string {
    return JSON.stringify(query);
  }

  applyTemplateVariables(query: NOAAQuery, scopedVars: ScopedVars): NOAAQuery {
    const templateSrv = getTemplateSrv();
    return {
      ...query,
      product: templateSrv.replace(query.product || '', scopedVars) as TCProduct,
      station: templateSrv.replace(query.station || '', scopedVars),
    };
  }
}
