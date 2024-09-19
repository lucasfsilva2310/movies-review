'use client'
export default function Example() {
  
    return (
      <header>

        <nav aria-label="Global" className="mx-auto flex max-w-full items-center justify-between p-6 lg:px-8">

          <div className="flex lg:flex-1">
            <a href="#" className="-m-1.5 p-1.5">
              <img alt="" src="https://cdn1.thepopcornfactory.com/wcsstore/ThePopcornFactory/images/catalog/tpf_85254x.jpg" className="h-8 w-auto" />
            </a>
          </div>

          <div className="lg:flex lg:flex-1 lg:justify-end">

            <a href="#" className="bg-slate-600 text-sm font-semibold leading-6 text-gray-900 rounded-md p-2 hover:bg-slate-700">
              Log in <span aria-hidden="true">&rarr;</span>
            </a>

          </div>

        </nav>

      </header>
    )
  }