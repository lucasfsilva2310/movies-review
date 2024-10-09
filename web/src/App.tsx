import { RouterProvider } from 'react-router-dom';
import Header from './components/Header';
import './index.css'
import { router } from './routes';

export const App = () => {

  return (
    <>
      <main className='bg-slate-500 w-screen h-screen'>
        <div className='w-full h-[10vh]'>
        <Header/>
        </div>
        <div className='w-full h-[90vh]'>
        <RouterProvider router={router}/>        
        </div>
      </main>
      
    </>
  )
}
