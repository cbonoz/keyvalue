import { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { apiClient, APIKey } from '../../api/client';
import { useSession } from '../../context/SessionContext';

export default function AppDetailPage() {
  const { id } = useParams<{ id: string }>();
  const [keys, setKeys] = useState<APIKey[]>([]);
  const [newKeyName, setNewKeyName] = useState('');
  const { session } = useSession();

  useEffect(() => {
    if (session && id) {
      apiClient.setSession(session);
      loadKeys();
    }
  }, [session, id]);

  const loadKeys = async () => {
    if (!id) return;
    const keys = await apiClient.listAppKeys(parseInt(id));
    setKeys(keys);
  };

  const handleCreateKey = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!newKeyName.trim() || !id) return;

    const key = await apiClient.createApiKey(parseInt(id), newKeyName);
    setNewKeyName('');
    loadKeys();

    // Show the new key to the user (only shown once)
    alert(`Your new API key is: ${key.key}\nPlease save it - you won't be able to see it again!`);
  };

  const handleDeleteKey = async (key: string) => {
    if (!id || !confirm('Are you sure you want to delete this API key?')) return;

    await apiClient.deleteApiKeys(parseInt(id), [key]);
    loadKeys();
  };

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">API Keys</h1>

      <form onSubmit={handleCreateKey} className="mb-6">
        <input
          type="text"
          value={newKeyName}
          onChange={(e) => setNewKeyName(e.target.value)}
          placeholder="API Key Name"
          className="p-2 border rounded mr-2"
        />
        <button type="submit" className="bg-blue-500 text-white px-4 py-2 rounded">
          Create API Key
        </button>
      </form>

      <div className="grid gap-4">
        {keys.map((key) => (
          <div key={key.id} className="border p-4 rounded flex justify-between items-center">
            <div>
              <h2 className="text-xl font-semibold">{key.name}</h2>
              <p className="text-gray-600">Created: {new Date(key.created_at).toLocaleDateString()}</p>
              {key.last_used && (
                <p className="text-gray-600">Last used: {new Date(key.last_used).toLocaleDateString()}</p>
              )}
            </div>
            <button
              onClick={() => handleDeleteKey(key.key)}
              className="bg-red-500 text-white px-4 py-2 rounded"
            >
              Delete
            </button>
          </div>
        ))}
      </div>
    </div>
  );
}
