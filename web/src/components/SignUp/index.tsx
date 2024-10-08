import { useCreateForm } from '../../hooks/useCreateForm';

interface ISignUpForm {
    username: string;
    password: string;
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
            {/* TODO: Change Div to correct element afterwards */}
            <div className="flex justify-center items-center w-full h-full">
                <div className="w-1/3 h-2/5 bg-slate-100 rounded-lg">
                    <form className="flex flex-col justify-center items-center w-full h-full" onSubmit={handleSubmit(handleSignUp)}>
                        <div className="flex flex-col w-4/5 h-1/5 mt-3 mb-2">
                            <label>
                                Username
                            </label>
                            <input {...assignInput('username', { required: true })} name="username"></input>
                        </div>

                        <div className="flex flex-col w-4/5 h-1/5 mt-3 mb-2">
                            <label>
                                Password
                            </label>
                            <input {...assignInput('password', { required: true })} name="password"></input>
                        </div>

                        <div>
                            <button className="bg-slate-300 p-3 rounded-md" type='submit'>Sign up</button>
                        </div>
                    </form>
                </div>
            </div>
        </>
    );
};