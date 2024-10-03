// import Header from './components/Header';
import './index.css'

function App() {

  return (
    <>
      <main className='bg-slate-500 w-screen h-screen'>
        {/* <Header/> */}

        {/* TODO: Change Div to correct element afterwards */}
        <div className='flex justify-center items-center w-full h-full'> {/* whole signUp page */}
            <div className='w-96 h-96 bg-slate-100 rounded-lg'> {/* whole signUp page */}
              <form className=''>
                <div>
                  <label>Username</label>
                  <input></input>
                </div>

                <div>
                  <label>Password</label>
                  <input></input>
                </div>
              </form>
            </div>
          </div>
      </main>
      
    </>
  )
}

export default App
