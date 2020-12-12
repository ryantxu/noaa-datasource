import { DataSourceInstanceSettings, ScopedVars } from '@grafana/data';
import { DataSourceWithBackend } from '@grafana/runtime';

import { NOAAQuery, NOAAOptions } from './types';

export class DataSource extends DataSourceWithBackend<NOAAQuery, NOAAOptions> {
  // Easy access for QueryEditor
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
    return 'TODO: ' + JSON.stringify(query);
  }

  applyTemplateVariables(query: NOAAQuery, scopedVars: ScopedVars): NOAAQuery {
    // if (!query.rawQuery) {
    //   return query;
    // }

    // const templateSrv = getTemplateSrv();
    // return {
    //   ...query,
    //   database: templateSrv.replace(query.database || '', scopedVars),
    //   table: templateSrv.replace(query.table || '', scopedVars),
    //   measure: templateSrv.replace(query.measure || '', scopedVars),
    //   rawQuery: templateSrv.replace(query.rawQuery), // DO NOT include scopedVars! it uses $__interval_ms!!!!!
    // };
    return query;
  }
}
