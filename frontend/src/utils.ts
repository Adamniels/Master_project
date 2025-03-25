export const getBackendIP = () => {
    return localStorage.getItem('backendIP') || 'http://localhost:8080'
}