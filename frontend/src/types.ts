export interface Admin1 {
  email: string;
  role: string;
}

export interface News {
  id: number;
  title: string;
  content: string;
  tag: string;
  date: string;
  created_at: string;
  updated_at: string;
}
export interface New1Response{
  news: News
}
export interface NewsResponse{
  news: News[]
}

export interface Services {
  id: number;
  title: string;
  description: string;
  price: number;
  category: string;
  icon_url: string;
  created_at: string;
  updated_at: string;
}
export interface Service1Response {
  service: Services;
}
export interface ServicesResponse{
  services: Services[]
}
export interface Team {
  id: number;
  name: string;
  position: string;
  experience: string;
  photo_url: string;
  created_at: string;
  updated_at: string;
}
export interface Team1Response{
  team_member: Team
}
export interface TeamResponse{
  team: Team[]
}
export interface Projects {
  id: number;
  title: string;
  description: string;
  category: string;
  status: string;
  created_at: string;
  updated_at: string;
}
export interface Project1Response{
  project: Projects
}
export interface ProjectsResponse{
  projects: Projects[]
}
export interface Statistics {
  violations_total: number;
  orders_total: number;
  fines_amount_total: number;
  collected_amount_total: number;
  evacuators_count: number;
  trips_count: number;
  evacuations_count: number;
  fine_lot_income: number;
  traffic_lights_active: number;
}
export interface StatisticsResponse{
  stats: Statistics
}
export interface Traffic {
  light_types: {
    t1: number;
    t2: number;
    t3: number;
  };
  install_years: {
    2020: number;
    2018: number;
    2015: number;
  };
}
export interface TrafficLight{
  id: number;
  adress: string;
  light_type: string;
  install_year: number;
  status: string;
  created_at: string;
  updated_at: string;
}
export interface TrafficLightsResponse{
  traffic_lights: TrafficLight[]
}

export interface Fine{
  id: string;
  date: string;
  violations_total: number;
  orders_total: number;
  fines_amount_total: number;
  collected_amount_total: number;
  created_at: string;
  updated_at: string;
}
export interface FinesResponse{
  fines: Fine[]
}

export interface Evacuation{
  id: number;
  date: string;
  evacuators_count: number;
  trips_count: number;
  evacuations_count: number;
  fine_lot_income: number;
  created_at: string;
  updated_at: string;
}
export interface EvacuationsResponse{
  evacuations: Evacuation[]
}

export interface EvacuationRoute{
  id: number;
  year: number;
  month: string;
  route: string;
  created_at: string;
  updated_at: string
}
export interface EvacuationRoutesResponse{
  evacuation_routes: EvacuationRoute[]
}