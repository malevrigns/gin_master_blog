export interface Author {
  id: number;
  username: string;
  email?: string;
  avatar?: string;
  bio?: string;
  role?: string;
}

export interface Category {
  id: number;
  name: string;
  slug: string;
  description?: string;
}

export interface Tag {
  id: number;
  name: string;
  slug: string;
}

export interface Article {
  id: number;
  title: string;
  slug: string;
  content?: string;
  excerpt?: string;
  cover_image?: string;
  views: number;
  likes: number;
  status: string;
  is_top: boolean;
  published_at?: string;
  created_at: string;
  updated_at: string;
  author?: Author;
  category?: Category;
  tags?: Tag[];
}

export interface LabHighlight {
  title: string;
  description?: string;
  tag?: string;
  category?: string;
  link?: string;
}

export interface LabResource {
  title: string;
  desc?: string;
  url: string;
  icon?: string;
}

export interface Lab {
  id: number;
  title: string;
  slug: string;
  subtitle?: string;
  badge?: string;
  badge_color?: string;
  description?: string;
  focus?: string;
  hero_image?: string;
  content?: string;
  highlights?: LabHighlight[] | null;
  resource_links?: LabResource[] | null;
}

export interface Link {
  id: number;
  name: string;
  url: string;
  logo?: string;
  desc?: string;
  is_visible: boolean;
  sort: number;
}

export interface Comment {
  id: number;
  article_id: number;
  content: string;
  author: string;
  email?: string;
  website?: string;
  parent_id?: number | null;
  status: string;
  created_at: string;
}

export interface Music {
  id: number;
  title: string;
  artist?: string;
  cover?: string;
  url: string;
  duration?: number;
  lrc?: string;
  is_public?: boolean;
}

export interface Playlist {
  id: number;
  name: string;
  description?: string;
  cover?: string;
}

export interface UserProfile {
  id: number;
  username: string;
  email?: string;
  role?: string;
  avatar?: string;
  bio?: string;
}

export interface PaginatedArticleResponse {
  articles: Article[];
  total: number;
  page: number;
  page_size: number;
}

export type RawLab = Omit<Lab, "highlights" | "resource_links"> & {
  highlights?: unknown;
  resource_links?: unknown;
};
