import { Session } from '@supabase/supabase-js';

export interface App {
  id: number;
  name: string;
  created_at: string;
  updated_at: string;
}

export interface APIKey {
  id: number;
  app_id: number;
  key: string;
  name: string;
  created_at: string;
  last_used?: string;
}

class APIClient {
  private baseUrl: string = '/api';
  private session: Session | null = null;

  setSession(session: Session | null) {
    this.session = session;
  }

  setBaseUrl(url: string) {
    this.baseUrl = url;
  }

  private async fetch(endpoint: string, options: RequestInit = {}) {
    const headers = {
      'Content-Type': 'application/json',
      ...(this.session && { Authorization: `Bearer ${this.session.access_token}` }),
      ...options.headers,
    };

    const response = await fetch(`${this.baseUrl}${endpoint}`, {
      ...options,
      headers,
    });

    if (!response.ok) {
      throw new Error(`API Error: ${response.statusText}`);
    }

    return response.json();
  }

  // Apps
  async listApps(): Promise<App[]> {
    return this.fetch('/apps');
  }

  async createApp(name: string): Promise<App> {
    return this.fetch('/apps', {
      method: 'POST',
      body: JSON.stringify({ name }),
    });
  }

  async deleteApp(id: number): Promise<void> {
    await this.fetch(`/apps/${id}`, { method: 'DELETE' });
  }

  // API Keys
  async listAppKeys(appId: number): Promise<APIKey[]> {
    return this.fetch(`/apps/${appId}/keys`);
  }

  async createApiKey(appId: number, name: string): Promise<APIKey> {
    return this.fetch(`/apps/${appId}/keys`, {
      method: 'POST',
      body: JSON.stringify({ name }),
    });
  }

  async deleteApiKeys(appId: number, keys: string[]): Promise<void> {
    await this.fetch(`/apps/${appId}/keys/delete`, {
      method: 'POST',
      body: JSON.stringify({ keys }),
    });
  }
}

export const apiClient = new APIClient();
apiClient.setBaseUrl(import.meta.env.VITE_SERVER_URL + '/api');
