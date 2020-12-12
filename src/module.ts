import { DataSourcePlugin } from '@grafana/data';
import { DataSource } from './DataSource';
import { NOAAQuery, NOAAOptions } from './types';
import { MetaInspector } from 'components/MetaInspector';
import { ConfigEditor } from 'components/ConfigEditor';
import { QueryEditor } from 'components/QueryEditor';

export const plugin = new DataSourcePlugin<DataSource, NOAAQuery, NOAAOptions>(DataSource)
  .setConfigEditor(ConfigEditor)
  .setMetadataInspector(MetaInspector)
  .setQueryEditor(QueryEditor);
