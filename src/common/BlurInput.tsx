import React, { PureComponent } from 'react';
import { Input } from '@grafana/ui';

interface Props {
  text?: string;
  placeholder?: string;
  onChange: (text: string) => void;
}

export class BlurInput extends PureComponent<Props> {
  onBlur = (e: React.FocusEvent<HTMLInputElement>) => {
    this.props.onChange(e.currentTarget.value);
  };

  onKeyPress = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter') {
      this.props.onChange(e.currentTarget.value);
    }
  };

  render() {
    const { text, placeholder } = this.props;
    return <Input defaultValue={text} onBlur={this.onBlur} onKeyPress={this.onKeyPress} placeholder={placeholder} />;
  }
}
