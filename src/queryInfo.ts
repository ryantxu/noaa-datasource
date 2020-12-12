import { SelectableValue } from '@grafana/data';
import { TCProduct } from 'types';

// water_level
//
//
// wind
// air_pressure	Barometric pressure as measured at the station.
// air_gap	Air Gap (distance between a bridge and the water's surface) at the station.
// conductivity	The water's conductivity as measured at the station.
// visibility	Visibility from the station's visibility sensor. A measure of atmospheric clarity.
// humidity	Relative humidity as measured at the station.
// salinity	Salinity and specific gravity data for the station.
// hourly_height	Verified hourly height water level data for the station.
// high_low	Verified high/low water level data for the station.
// daily_mean	Verified daily mean water level data for the station.
// monthly_mean	Verified monthly mean water level data for the station.
// one_minute_water_level	One minute water level data for the station.
//
// datums	datums data for the stations.
// currents	Currents data for currents stations.
// currents_predictions	Currents predictions data for currents predictions stations.

export const tidesAndCurrentsProducts: Array<SelectableValue<TCProduct>> = [
  {
    label: 'Water Level',
    value: TCProduct.water_level,
    description: 'Preliminary or verified water levels, depending on availability.',
  },
  {
    label: 'Air Temperature',
    value: TCProduct.air_temperature,
    description: 'Air temperature as measured at the station.',
  },
  {
    label: 'Water Temperature',
    value: TCProduct.water_temperature,
    description: 'Water temperature as measured at the station.',
  },
  { label: 'Wind', value: TCProduct.wind, description: 'Wind speed, direction, and gusts as measured at the station.' },
  {
    label: 'Predictions',
    value: TCProduct.predictions,
    description: '6 minute predictions water level data for the station.*',
  },
  //   { label: '', value: TCProduct.water_level, description: '' },
  //   { label: '', value: TCProduct.water_level, description: '' },
  //   { label: '', value: TCProduct.water_level, description: '' },
  //   { label: '', value: TCProduct.water_level, description: '' },
  //   { label: '', value: TCProduct.water_level, description: '' },
  //   { label: '', value: TCProduct.water_level, description: '' },
];
