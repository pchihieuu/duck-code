export interface User {
  id: string;
  name: string;
  email: string;
  avatarUrl?: string;
  role: 'student' | 'admin';
  level: number;
  xp: number;
  streakDays: number;
}
