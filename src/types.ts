import { DataSourceJsonData } from '@grafana/data';
import { DataQuery } from '@grafana/schema';

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
  monthly_mean = 'monthly_mean',
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

// Matches https://github.com/ryantxu/noaa-datasource/blob/main/pkg/models/query.go#L3
export enum QueryDate {
  Query = 'query',
  Today = 'today',
  Recent = 'recent',
  Latest = 'latest',
}

// date=today
// The last 3 days of data
// date=recent
// The last data point available within the last 18 min
// date=latest

export interface NOAAQuery extends DataQuery {
  queryType?: QueryType;
  station?: string;
  product?: TCProduct;
  units?: 'metric' | 'english';
  date?: QueryDate;
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
