using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.Options;

namespace ReloadConfigurationSample.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class ValuesController : ControllerBase
    {
        private readonly StaticOptions _staticConfig;
        private readonly ConfigMapOptions _mapConfig;
        public ValuesController(IOptionsSnapshot<StaticOptions> staticConfig, IOptionsSnapshot<ConfigMapOptions> mapConfig)
        {
            _staticConfig = staticConfig.Value;
            _mapConfig = mapConfig.Value;
        }
        // GET api/values
        [HttpGet]
        public ActionResult<IEnumerable<string>> Get()
        {
            return new string[] { Environment.MachineName, _mapConfig.Hello, _staticConfig.Version };
        }
    }
}
