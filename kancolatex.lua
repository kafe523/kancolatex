local kancolatex = {}

local ffi = require("ffi")

--- split a string
--- @param s string
--- @param delimiter string
--- @return table
local function strSplit(s, delimiter)
    -- https://gist.github.com/jaredallard/ddb152179831dd23b230
    local result = { }
    local from  = 1
    local delim_from, delim_to = string.find( s, delimiter, from  )
    while delim_from do
        table.insert( result, string.sub( s, from , delim_from-1 ) )
        from  = delim_to + 1
        delim_from, delim_to = string.find( s, delimiter, from  )
    end
    table.insert( result, string.sub( s, from  ) )
    return result
end

--- On windows is `libkancolatex.dll`, OSX is `libkancolatex.dylib`, and linux is `libkancolatex.so`.
---@return string
function kancolatex.GenerateDefaultLibName()
    local os = ffi.os
    local arch = ffi.arch

    if os == "Windows" then
        return "libkancolatex.dll"
    elseif os == "OSX" then
        return "libkancolatex.dylib"
    elseif (os == "Linux" or os == "BSD" or os == "POSIX") then
        return "libkancolatex.so"
    end

    return "libkancolatex-"..os.."-"..arch..".unsupported"
end

ffi.cdef[[
uintptr_t InitMacroContext(char* cNoroPath);
char* AccessMacro(uintptr_t mcp, char* mac);
char* AccessPattern(uintptr_t mcp, char* pat);
char* AccessKeyStr(uintptr_t mcp);
void Free(void* ptr);
size_t StrLen(char* sp);
]]
--- Shared Library singleton, Should manual init with InitLoadLib
local libkancolatex;

---Convert char* to lua string
---@param sp number
---@return string
local function ccToStr(sp)
    local s = ffi.string(sp, ffi.cast("unsigned long long", libkancolatex.StrLen(sp)))
    libkancolatex.Free(sp)
    return s
end

--- Convert lua string to char*
---@param s string
---@return number
local function strToCc(s)
    local cs = ffi.cast("char *", ffi.new("const char *", s))
    return cs
end

---Specific libkancolatex library path for the kancolatex package.
---@param libPath string
function kancolatex.InitLoadLib(libPath)
    libkancolatex = ffi.load(libPath)
    if libkancolatex == nil then
        error("libkancolatex failed to init" .. type(libkancolatex))
    end
end

---Init Kancolatex Macro Context Only, Load noro json dump from path.
---@param noroPath string path to noro json dump,
---@return number pointer to macro context.
function kancolatex.InitContext(noroPath)
    return libkancolatex.InitMacroContext(strToCc(noroPath))
end

---@param kctxp number pointer to macro context.
---@return string
local function accessKeyString(kctxp)
    local cc = libkancolatex.AccessKeyStr(kctxp)
    return ccToStr(cc)
end

---Init Kancolatex predefined macro
---@param kctxp number pointer to macro context.
function kancolatex.InitLuaTeXMacro(kctxp)
    for _, value in ipairs(strSplit(accessKeyString(kctxp), ",")) do
        local result = ccToStr(libkancolatex.AccessMacro(kctxp, strToCc(value)))
        token.set_macro(value, result)
    end

end

---Getting string result from Kancolatex Macro Context with special access pattern.
---Indexing array are also chaining with dot. kancolatex does not implement [].
---
---Example:
---
---`CalcManager.FleetInfo.Fleets.0.Ships.0.Data.Name` will retrieve Ship name in Japanese.
---
---@param kctxp number pointer to macro context.
---@param pat string access pattern.
---@return string _ no result or nil will return empty string
function kancolatex.AccessPattern(kctxp, pat)
    local cc = libkancolatex.AccessPattern(ffi.cast("int", kctxp), strToCc(pat))
    return ccToStr(cc)
end

---Getting string result from Kancolatex Macro Context with predefined macro.
---@param kctxp number pointer to macro context.
---@param macro string macro without `\`
---@return string no result or nil will return empty string
function kancolatex.AccessMacro(kctxp, macro)
    local cc = libkancolatex.AccessMacro(ffi.cast("int", kctxp), strToCc(macro))
    return ccToStr(cc)
end

return kancolatex
