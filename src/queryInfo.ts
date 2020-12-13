import { SelectableValue } from '@grafana/data';
import { QueryDate, TCProduct } from 'types';

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
    label: 'High Low',
    value: TCProduct.high_low,
    description: 'Verified high/low water level data for the station.',
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
  {
    label: 'Air pressure',
    value: TCProduct.air_pressure,
    description: 'Barometric pressure as measured at the station.',
  },
  {
    label: 'air_gap',
    value: TCProduct.air_gap,
    description: `Air Gap (distance between a bridge and the water's surface) at the station.`,
  },
  {
    label: 'conductivity',
    value: TCProduct.conductivity,
    description: `The water's conductivity as measured at the station.`,
  },
  {
    label: 'visibility',
    value: TCProduct.visibility,
    description: `Visibility from the station's visibility sensor. A measure of atmospheric clarity.`,
  },
  { label: 'humidity', value: TCProduct.humidity, description: `` },
  { label: 'salinity', value: TCProduct.salinity, description: `Relative humidity as measured at the station.` },
  {
    label: 'hourly_height',
    value: TCProduct.hourly_height,
    description: `Verified hourly height water level data for the station.`,
  },
  {
    label: 'daily_mean',
    value: TCProduct.daily_mean,
    description: `Verified daily mean water level data for the station.`,
  },
  {
    label: 'monthly_mean',
    value: TCProduct.monthly_mean,
    description: `Verified monthly mean water level data for the station.`,
  },
  {
    label: 'one_minute_water_level',
    value: TCProduct.one_minute_water_level,
    description: `One minute water level data for the station.`,
  },
];

export const dateOptions: Array<SelectableValue<QueryDate>> = [
  {
    label: 'From query',
    value: QueryDate.Query,
    description: 'Use the dashboard query to get the date',
  },
  {
    label: 'Today',
    value: QueryDate.Today,
    description: 'Todays reading (in local time)',
  },
  {
    label: 'Recent',
    value: QueryDate.Recent,
    description: 'The last 3 days of data',
  },
  {
    label: 'Latest',
    value: QueryDate.Latest,
    description: 'The last data point available within the last 18 min',
  },
];
