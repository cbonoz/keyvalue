import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { apiClient, App } from '../../api/client';
import { useSession } from '../../context/SessionContext';

export default function AppsPage() {
  const [apps, setApps] = useState<App[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [newAppName, setNewAppName] = useState('');
  const { session } = useSession();

  useEffect(() => {
    if (session) {
      apiClient.setSession(session);
    }
      loadApps();
  }, [session]);

  const loadApps = async () => {
    if (loading ||!session) {return;}
    try {
      setLoading(true);
    const apps = await apiClient.listApps();
    setApps(apps);
    } catch (error) {
      console.error('Error loading apps:', error);
      alert('Failed to load apps. Please try again later.');
    } finally {
      setLoading(false);
    }
  };

  const handleCreateApp = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!newAppName.trim()) return;

    await apiClient.createApp(newAppName);
    setNewAppName('');
    loadApps();
  };

  const handleDeleteApp = async (id: number) => {
    if (confirm('Are you sure you want to delete this app?')) {
      await apiClient.deleteApp(id);
      loadApps();
    }
  };

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">Your Apps</h1>

      <form onSubmit={handleCreateApp} className="mb-6">
        <input
          type="text"
          value={newAppName}
          onChange={(e) => setNewAppName(e.target.value)}
          placeholder="New App Name"
          className="p-2 border rounded mr-2"
        />
        <button type="submit" className="bg-blue-500 text-white px-4 py-2 rounded">
          Create App
        </button>
      </form>

      <div className="grid gap-4">
        {apps.map((app) => (
          <div key={app.id} className="border p-4 rounded flex justify-between items-center">
            <div>
              <h2 className="text-xl font-semibold">{app.name}</h2>
              <p className="text-gray-600">Created: {new Date(app.created_at).toLocaleDateString()}</p>
            </div>
            <div>
              <Link
                to={`/apps/${app.id}`}
                className="bg-green-500 text-white px-4 py-2 rounded mr-2"
              >
                Manage
              </Link>
              <button
                onClick={() => handleDeleteApp(app.id)}
                className="bg-red-500 text-white px-4 py-2 rounded"
              >
                Delete
              </button>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
