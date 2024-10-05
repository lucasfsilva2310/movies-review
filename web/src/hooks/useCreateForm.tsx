import {
    FieldValues,
    useForm,
} from 'react-hook-form';

export const useCreateForm = <T extends FieldValues,>() => {

    const { register, handleSubmit } = useForm<T>()

    return {
        assignInput: register,
        handleSubmit
    }
}