When create a middleware global handle error:
    1/ Collect all error and return a JSON format in one location
    --> Question: what location should place global handle error ?
    --> Answer: In a middleware layer because:
        + in the data flow: request: request -> route -> middleware -> controller -> service -> repo 
                            respone: repo -> service -> controller -> middleware -> respone

        + so all error will gather together in controller layer and middleware always run first in request and last in respone --> Put global handle error in middleware.
    
    2/ Implement global error
    --> Question: What is input and output ?
    --> Answer: Input error and Output json format with error of not.

    --> Question: How to input in Gin ?
    --> Answer: Use Error() in gin.Context to fetch an error.

    --> Question: How to modify the output ?
    --> Answer: 
            1/ Definition the error.
            2/ Implement logic in function global handle error.

Flow data:
    1/ Create Default Engin:        server := gin.Default()

    2/ Register Middleware Handler: server.Use(Errorhandler())
        Note: 
            The Handler fuction accepts the request and return the respone
            In Gin, it define the HandlerFunc (Handler function) that have the form:
            type HandlerFunc func(*Context)

            The Use() is a variadic function that take in the slice HandlerFunc and return the IRoutes:
            func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes {
                engine.RouterGroup.Use(middleware...)
                engine.rebuild404Handlers()
                engine.rebuild405Handlers()
                return engine
            }

    3/ Call the Use() in RouterGroup struct: 
        // Use adds middleware to the group
        func (group *RouterGroup) Use(middleware ...HandlerFunc) IRoutes {
            group.Handlers = append(group.Handlers, middleware...)
            return group.returnObj()
        }

        // IRouter defines all router handle interface includes single and group router.
        type IRouter interface {
            IRoutes
            Group(string, ...HandlerFunc) *RouterGroup
        }

        // IRoutes defines all router handle interface.
        type IRoutes interface {
            Use(...HandlerFunc) IRoutes

            Handle(string, string, ...HandlerFunc) IRoutes
            Any(string, ...HandlerFunc) IRoutes
            GET(string, ...HandlerFunc) IRoutes
            POST(string, ...HandlerFunc) IRoutes
            DELETE(string, ...HandlerFunc) IRoutes
            PATCH(string, ...HandlerFunc) IRoutes
            PUT(string, ...HandlerFunc) IRoutes
            OPTIONS(string, ...HandlerFunc) IRoutes
            HEAD(string, ...HandlerFunc) IRoutes
            Match([]string, string, ...HandlerFunc) IRoutes

            StaticFile(string, string) IRoutes
            StaticFileFS(string, string, http.FileSystem) IRoutes
            Static(string, string) IRoutes
            StaticFS(string, http.FileSystem) IRoutes
        }
        
        // RouterGroup is used internally to configure router, a RouterGroup is associated with
        // a prefix and an array of handlers (middleware).
        type RouterGroup struct {
            Handlers HandlersChain
            basePath string
            engine   *Engine
            root     bool
        }

        The Handlers field has the HandlersChain type:
        // HandlersChain defines a HandlerFunc slice.
        type HandlersChain []HandlerFunc

        The return Obj() will return the obj route which contains the engin field and the Handlers field:
        func (group *RouterGroup) returnObj() IRoutes {
            if group.root {
                return group.engine
            }
            return group
        }

