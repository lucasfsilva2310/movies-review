import { createBrowserRouter } from 'react-router-dom';
import { SignUp } from '../components/SignUp';
import { LandingPage } from '../components/LandingPage';
import { UnexpectedError } from '../components/UnexpectedError';

export const router = createBrowserRouter([
    {
        path: '/',
        element: <LandingPage/>,
        errorElement: <UnexpectedError/>
    },
    {
        path: '/sign-up',
        element: <SignUp/>
    }
])