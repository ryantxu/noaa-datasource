import { DataQuery, DataSourceJsonData } from '@grafana/data';

export enum TCProduct {
  water_level = 'water_level',
  air_temperature = 'air_temperature',
  water_temperature = 'water_temperature',
  wind = 'wind',
  air_pressure = 'air_pressure',
  air_gap = 'air_gap',
  conductivity = 'conductivity',
  visibility = 'visibility',
  humidity = 'humidity',
  salinity = 'salinity',
  hourly_height = 'hourly_height',
  high_low = 'high_low',
  daily_mean = 'daily_mean',
  one_minute_water_level = 'one_minute_water_level',
  predictions = 'predictions',
  datums = 'datums',
  currents = 'currents',
  currents_predictions = 'currents_predictions',
}

// Matches https://github.com/ryantxu/noaa-datasource/blob/main/pkg/models/query.go#L3
export enum QueryType {
  TidesAndCurrents = 'TidesAndCurrents',
}

export interface NOAAQuery extends DataQuery {
  queryType?: QueryType;
  station?: number;
  product?: TCProduct;
  units?: 'metric' | 'english';
}

/**
 * Metadata attached to DataFrame results
 */
export interface NOAACustomMeta {
  id?: number;
  name?: string;
  lat?: number;
  lon?: number;
}

/**
 * Global datasource options
 */
export interface NOAAOptions extends DataSourceJsonData {}
