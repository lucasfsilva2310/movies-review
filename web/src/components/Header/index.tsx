'use client';
export default function Header() {
    return (
        <header className="w-full h-20 bg-slate-500">
            <nav
                aria-label="Global"
                className={`flex items-center h-full justify-between bg-slate-600 p-2`}
            >
                <div className="flex lg:flex-1">
                    <a
                        href="/"
                        className="h-16 w-16 object-contain"
                    >
                        <img
                            alt=""
                            src="https://cdn1.thepopcornfactory.com/wcsstore/ThePopcornFactory/images/catalog/tpf_85254x.jpg"
                            className={`h-full w-full`}
                        />
                    </a>
                </div>

                <div className="flex">
                    <a
                        href="/movies"
                        className={`rounded-md bg-slate-700 p-2 text-sm font-semibold  text-gray-900 hover:bg-slate-800`}
                    >
                        Movies{' '}
                    </a>
                </div>

                <div className="lg:flex lg:flex-1 lg:justify-end">
                    <a
                        href="#"
                        className={`rounded-md bg-slate-700 p-2 text-sm font-semibold  text-gray-900 transition ease-in-ou hover:bg-slate-800`}
                    >
                        Log in{' '}
                        <span aria-hidden="true">
                            &rarr;
                        </span>
                    </a>
                </div>
            </nav>
        </header>
    );
}