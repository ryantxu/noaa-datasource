import React, { PureComponent } from 'react';
import { DataSourcePluginOptionsEditorProps } from '@grafana/data';
import { NOAAOptions } from '../types';

export type Props = DataSourcePluginOptionsEditorProps<NOAAOptions>;

export class ConfigEditor extends PureComponent<Props> {
  render() {
    return (
      <div>
        See: <a href="https://tidesandcurrents.noaa.gov/">https://tidesandcurrents.noaa.gov/</a>
      </div>
    );
  }
}
