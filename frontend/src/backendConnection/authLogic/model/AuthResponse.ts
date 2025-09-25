import type { Admin1 } from "../../types";
 export interface AuthResponse{
    accessToken: string;
    refreshToken: string;
    admin: Admin1;
 }