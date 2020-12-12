import React, { PureComponent } from 'react';
import { QueryEditorProps, SelectableValue } from '@grafana/data';
import { DataSource } from '../DataSource';
import { NOAAQuery, NOAAOptions, TCProduct } from '../types';
import { InlineField, Select } from '@grafana/ui';
import { tidesAndCurrentsProducts } from '../queryInfo';
import { BlurInput } from 'common/BlurInput';

type Props = QueryEditorProps<DataSource, NOAAQuery, NOAAOptions>;

const labelWidth = 10;

export class QueryEditor extends PureComponent<Props> {
  onProductChange = (sel: SelectableValue<TCProduct>) => {
    const { onChange, query, onRunQuery } = this.props;
    onChange({ ...query, product: sel.value! });
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

          <InlineField label="Station" labelWidth={labelWidth} grow={true}>
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
