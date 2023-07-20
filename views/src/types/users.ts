
export interface User {
  id: number;
  username: string;
  first_name: string;
  last_name: string;
  email?: string;
  phone_number?: string;
  role: string;
}

interface UpdateUserBody {
  username?: string;
  first_name?: string;
  last_name?: string;
  email?: string;
  phone_number?: string;
  role?: string;
}
export interface UpdateUser {
  user_id: number | string;
  body: UpdateUserBody;
}
