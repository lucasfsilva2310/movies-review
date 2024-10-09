import { createBrowserRouter } from 'react-router-dom';
import { SignUp } from '../components/SignUp';
import { LandingPage } from '../components/LandingPage';

export const router = createBrowserRouter([
    {
        path: '/',
        element: <LandingPage/>
    },
    {
        path: '/sign-up',
        element: <SignUp/>
    }
])