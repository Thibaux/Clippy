declare module '*.module.css';
declare module '*.module.scss';

type Item = {
  id: number;
  content: string;
  filters: [{ id: number; name: string }];
  createdAt: string;
  updatedAt: string;
};
