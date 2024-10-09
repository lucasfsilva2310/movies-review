import { useCreateForm } from '../../hooks/useCreateForm';

interface ISignUpForm {
    username: string;
    email: string;
    password: string;
    repeatPassword: string;
}

export const SignUp = () => {
    const {
        assignInput,
        handleSubmit
    } = useCreateForm<ISignUpForm>()

    const handleSignUp = (data: ISignUpForm) => {
        console.log(data)
    }

    return (
        <>
            <div className="flex justify-center items-center w-full h-full">
                <div className="w-1/3 h-5/6 bg-slate-100 rounded-lg max-[1020px]:w-full max-[1020px]:h-full">
                    <form className="flex flex-col justify-evenly items-center w-full h-full" onSubmit={handleSubmit(handleSignUp)}>
                        <div className="flex flex-col w-4/5 h-1/5 mt-12 mb-2">
                            <label className="mb-2" htmlFor='username'>
                                Username
                            </label>
                            <input className='p-3' {...assignInput('username', { required: true })} name="username"></input>
                        </div>

                        <div className="flex flex-col w-4/5 h-1/5 mt-3 mb-2">
                            <label className="mb-2" htmlFor='email'>
                                email
                            </label>
                            <input className='p-3' {...assignInput('email', { required: true })} name="email"></input>
                        </div>

                        <div className="flex flex-col w-4/5 h-1/5 mt-3 mb-2">
                            <label className="mb-2" htmlFor='password'>
                                Password
                            </label>
                            <input className='p-3' {...assignInput('password', { required: true })} name="password"></input>
                        </div>

                        <div className="flex flex-col w-4/5 h-1/5 mt-3 mb-12">
                            <label className="mb-2" htmlFor='repeatPassword'>
                                Repeat password
                            </label>
                            <input className='p-3' {...assignInput('repeatPassword', { required: true })} name="password"></input>
                        </div>

                        <div>
                            <button className="bg-slate-300 mb-20 p-3 rounded-md transition ease-in-out hover:bg-slate-400" type='submit'>Sign up</button>
                        </div>
                    </form>
                </div>
            </div>
        </>
    );
};