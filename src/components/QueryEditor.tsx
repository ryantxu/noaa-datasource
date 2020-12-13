import React, { PureComponent } from 'react';
import { QueryEditorProps, SelectableValue } from '@grafana/data';
import { DataSource } from '../DataSource';
import { NOAAQuery, NOAAOptions, TCProduct, QueryDate } from '../types';
import { InlineField, Select } from '@grafana/ui';
import { dateOptions, tidesAndCurrentsProducts } from '../queryInfo';
import { BlurInput } from 'common/BlurInput';

type Props = QueryEditorProps<DataSource, NOAAQuery, NOAAOptions>;

export const units: Array<SelectableValue<string>> = [
  {
    label: 'English',
    value: 'english',
  },
  {
    label: 'Metric',
    value: 'metric',
  },
];

const labelWidth = 10;

export class QueryEditor extends PureComponent<Props> {
  onProductChange = (sel: SelectableValue<TCProduct>) => {
    const { onChange, query, onRunQuery } = this.props;
    onChange({ ...query, product: sel.value! });
    onRunQuery();
  };

  onUnitsChange = (sel: SelectableValue<string>) => {
    const { onChange, query, onRunQuery } = this.props;
    onChange({ ...query, units: sel.value as any });
    onRunQuery();
  };

  onDateChange = (sel: SelectableValue<QueryDate>) => {
    const { onChange, query, onRunQuery } = this.props;
    onChange({ ...query, date: sel.value as any });
    onRunQuery();
  };

  onStationChange = (txt: string) => {
    const { onChange, query, onRunQuery } = this.props;
    onChange({ ...query, station: +txt });
    onRunQuery();
  };

  render() {
    const { query } = this.props;

    return (
      <>
        <div className="gf-form">
          <InlineField label="Product" labelWidth={labelWidth} grow={true}>
            <Select
              options={tidesAndCurrentsProducts}
              value={tidesAndCurrentsProducts.find(v => v.value === query.product)}
              onChange={this.onProductChange}
              placeholder="Select query type"
              menuPlacement="bottom"
            />
          </InlineField>
          <InlineField label="Unit">
            <Select
              options={units}
              value={units.find(v => v.value === query.units) || units[0]}
              onChange={this.onUnitsChange}
              placeholder="Select units"
              menuPlacement="bottom"
            />
          </InlineField>
          <InlineField label="Date">
            <Select
              options={dateOptions}
              value={dateOptions.find(v => v.value === query.date) || dateOptions[0]}
              onChange={this.onDateChange}
              placeholder="Select date"
              menuPlacement="bottom"
            />
          </InlineField>
          <InlineField label="Station" grow={true}>
            <BlurInput
              text={query.station ? `${query.station}` : ''}
              placeholder="Station ID"
              onChange={this.onStationChange}
            />
          </InlineField>
        </div>
      </>
    );
  }
}
